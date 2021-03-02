package uci

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/rotisserie/eris"
)

// Options, for initializing the chess engine
type Options struct {
	MultiPV int  // number of principal variations (ranks top X moves)
	Hash    int  // hash size in MB
	Ponder  bool // whether the engine should ponder
	OwnBook bool // whether the engine should use its opening book
	Threads int  // max number of threads the engine should use
}

// scoreKey helps us save the latest unique result where unique is
// defined as having unique values for each of the fields
type scoreKey struct {
	Depth   int
	MultiPV int
}

// ScoreResult holds the score result records returned
// by the engine
type ScoreResult struct {
	Time           int      // time spent to get this result (ms)
	Depth          int      // depth (number of plies) of result record
	SelDepth       int      // selective depth -- some engines don't report this
	Nodes          int      // total nodes searched to get this result
	NodesPerSecond int      // current nodes per second rate
	MultiPV        int      // 0 if MultiPV not set
	Score          int      // score centipawns or mate in X if Mate is true
	Mate           bool     // whether this move results in forced mate
	BestMoves      []string // best line for this result
}

// Results holds a slice of ScoreResult records
// as well as some overall result data
type Results struct {
	BestMove string
	results  map[scoreKey]ScoreResult
	Results  []ScoreResult
}

func (r Results) String() string {
	b, _ := json.MarshalIndent(r, "", "  ")
	return fmt.Sprintln(string(b))
}

// Engine holds the information needed to communicate with
// a chess engine executable. Engines should be created with
// a call to NewEngine(/path/to/executable)
type Engine struct {
	cmd    *exec.Cmd
	stdout *bufio.Reader
	stdin  *bufio.Writer
}

// NewEngine returns an Engine it has spun up
// and connected communication to
func NewEngine(path string, arg ...string) (*Engine, error) {
	eng := Engine{}
	eng.cmd = exec.Command(path, arg...)
	stdin, err := eng.cmd.StdinPipe()
	if err != nil {
		return nil, eris.Wrap(err, "")
	}

	stdout, err := eng.cmd.StdoutPipe()
	if err != nil {
		return nil, eris.Wrap(err, "")
	}

	if err := eng.cmd.Start(); err != nil {
		return nil, eris.Wrap(err, "")
	}

	eng.stdin = bufio.NewWriter(stdin)
	eng.stdout = bufio.NewReader(stdout)

	return &eng, nil
}

// SetOptions sends setoption commands to the Engine
// for the values set in the Options record passed in
func (eng *Engine) SetOptions(opt Options) error {
	var err error
	if opt.MultiPV > 0 {
		err = eng.SendOption("multipv", opt.MultiPV)
		if err != nil {
			return eris.Wrap(err, "")
		}
	}

	if opt.Hash > 0 {
		err = eng.SendOption("hash", opt.Hash)
		if err != nil {
			return err
		}
	}

	if opt.Threads > 0 {
		err = eng.SendOption("threads", opt.Threads)
		if err != nil {
			return eris.Wrap(err, "")
		}
	}

	err = eng.SendOption("ownbook", opt.OwnBook)
	if err != nil {
		return eris.Wrap(err, "")
	}

	err = eng.SendOption("ponder", opt.Ponder)
	if err != nil {
		return eris.Wrap(err, "")
	}

	return eris.Wrap(err, "")
}

// SendOption sends setoption command to the Engine
func (eng *Engine) SendOption(name string, value interface{}) error {
	_, err := eng.stdin.WriteString(
		fmt.Sprintf("setoption name %s value %v\n", name, value),
	)
	if err != nil {
		return eris.Wrap(err, "")
	}

	err = eng.stdin.Flush()
	return eris.Wrap(err, "")
}

// SetFEN takes a FEN string and tells the engine to set the position
func (eng *Engine) SetFEN(fen string) error {
	_, err := eng.stdin.WriteString(
		fmt.Sprintf("position fen %s\n", fen),
	)
	if err != nil {
		return eris.Wrap(err, "")
	}

	err = eng.stdin.Flush()
	return eris.Wrap(err, "")
}

func (eng *Engine) SetMoves(moves string) error {
	_, err := eng.stdin.WriteString(
		fmt.Sprintf("position startpos moves %s\n", moves),
	)
	if err != nil {
		return eris.Wrap(err, "")
	}

	err = eng.stdin.Flush()
	return eris.Wrap(err, "")
}

// Go can use search moves, depth and time to move as filter  for the results being returned.
// see http://wbec-ridderkerk.nl/html/UCIProtocol.html
func (eng *Engine) Go(depth int, searchmoves string, movetime int64) (*Results, error) {
	res := Results{}
	goCmd := "go "

	if depth != 0 {
		goCmd += fmt.Sprintf("depth %d", depth)
	}

	if searchmoves != "" {
		goCmd += fmt.Sprintf(" searchmoves %s", searchmoves)
	}

	if movetime != 0 {
		goCmd += fmt.Sprintf(" movetime %d", movetime)
	}

	goCmd += "\n"

	_, err := eng.stdin.WriteString(goCmd)
	if err != nil {
		return nil, eris.Wrap(err, "")
	}

	err = eng.stdin.Flush()
	if err != nil {
		return nil, eris.Wrap(err, "")
	}

	for {
		line, err := eng.stdout.ReadString('\n')
		if err != nil {
			return nil, eris.Wrap(err, "")
		}

		line = strings.Trim(line, "\n")
		if strings.HasPrefix(line, "bestmove") {
			dummy := ""

			_, err := fmt.Sscanf(line, "%s %s", &dummy, &res.BestMove)
			if err != nil {
				return nil, eris.Wrap(err, "")
			}

			break
		}

		err = res.addLineToResults(line)
		if err != nil {
			return nil, eris.Wrap(err, "")
		}
	}

	for _, v := range res.results {
		if v.Depth != depth {
			continue
		}

		res.Results = append(res.Results, v)
	}

	sort.Sort(byDepth(res.Results))

	return &res, nil
}

// GoDepth takes a depth and an optional uint flag that configures filters
// for the results being returned.
func (eng *Engine) GoDepth(depth int, resultOpts ...uint) (*Results, error) {
	return eng.Go(depth, "", 0)
}

type byDepth []ScoreResult

func (a byDepth) Len() int      { return len(a) }
func (a byDepth) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byDepth) Less(i, j int) bool {
	if a[i].Depth == a[j].Depth {
		if a[i].MultiPV == a[j].MultiPV {
		}

		return a[i].MultiPV < a[j].MultiPV
	}
	return a[i].Depth < a[j].Depth
}

func (res *Results) addLineToResults(line string) error {
	var err error
	if !strings.HasPrefix(line, "info") {
		return nil
	}

	log.Println(line)

	rd := strings.NewReader(line)
	s := scanner.Scanner{}
	s.Init(rd)
	s.Mode = scanner.ScanIdents | scanner.ScanChars | scanner.ScanInts
	r := ScoreResult{}

	for s.Scan() != scanner.EOF {
		switch s.TokenText() {
		case "info":
		case "currmove":
			return nil
		case "depth":
			s.Scan()
			r.Depth, err = strconv.Atoi(s.TokenText())
			if err != nil {
				return eris.Wrap(err, "")
			}
		case "seldepth":
			s.Scan()
			r.SelDepth, err = strconv.Atoi(s.TokenText())
			if err != nil {
				return eris.Wrap(err, "")
			}
		case "time":
			s.Scan()
			r.Time, err = strconv.Atoi(s.TokenText())
			if err != nil {
				return eris.Wrap(err, "")
			}
		case "nodes":
			s.Scan()
			r.Nodes, err = strconv.Atoi(s.TokenText())
			if err != nil {
				return eris.Wrap(err, "")
			}
		case "nps":
			s.Scan()
			r.NodesPerSecond, err = strconv.Atoi(s.TokenText())
			if err != nil {
				return eris.Wrap(err, "")
			}
		case "multipv":
			s.Scan()
			r.MultiPV, err = strconv.Atoi(s.TokenText())
			if err != nil {
				return eris.Wrap(err, "")
			}
		case "score":
			s.Scan()
			switch s.TokenText() {
			case "cp":
				s.Scan()
			case "mate":
				r.Mate = true
				s.Scan()
			}
			negative := 1
			if s.TokenText() == "-" {
				negative = -1
				s.Scan()
			}
			r.Score, err = strconv.Atoi(s.TokenText())
			if err != nil {
				return eris.Wrap(err, "")
			}
			r.Score = r.Score * negative
		case "pv":
			for s.Scan() != scanner.EOF {
				r.BestMoves = append(r.BestMoves, s.TokenText())
			}
		}
	}

	if r.Depth > 0 {
		if res.results == nil {
			res.results = make(map[scoreKey]ScoreResult)
		}
		res.results[scoreKey{
			Depth:   r.Depth,
			MultiPV: r.MultiPV,
		}] = r
	}
	return nil
}

func (eng *Engine) Close() {
	_, err := eng.stdin.WriteString("stop\n")
	if err != nil {
		log.Println("failed to stop engine: ", eris.ToString(err, true))
	}

	err = eng.stdin.Flush()
	if err != nil {
		log.Println("failed to stop engine: ", eris.ToString(err, true))
	}

	err = eng.cmd.Process.Kill()
	if err != nil {
		log.Println("failed to stop engine: ", eris.ToString(err, true))
	}

	err = eng.cmd.Wait()
	if err != nil {
		log.Println("failed to stop engine: ", eris.ToString(err, true))
	}
}
