package handler

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/klauspost/cpuid/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

func main() {
	fmt.Println("Name:", CPU.BrandName)
	fmt.Println("PhysicalCores:", CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", CPU.LogicalCores)
	fmt.Println("Family", CPU.Family, "Model:", CPU.Model, "Vendor ID:", CPU.VendorID)
	fmt.Println("Features:", fmt.Sprintf(strings.Join(CPU.FeatureSet(), ",")))
	fmt.Println("Cacheline bytes:", CPU.CacheLine)
	fmt.Println("L1 Data Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", CPU.Hz, "hz")

	// Test if we have these specific features:
	if CPU.Supports(SSE, SSE2) {
		fmt.Println("We have Streaming SIMD 2 Extensions")
	}
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/calculate", handlers.Calculate)

	e.Logger.Fatal(e.Start(":1323"))
}

// Handler is required by Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Name:", CPU.BrandName)
	fmt.Println("PhysicalCores:", CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", CPU.LogicalCores)
	fmt.Println("Family", CPU.Family, "Model:", CPU.Model, "Vendor ID:", CPU.VendorID)
	fmt.Println("Features:", fmt.Sprintf(strings.Join(CPU.FeatureSet(), ",")))
	fmt.Println("Cacheline bytes:", CPU.CacheLine)
	fmt.Println("L1 Data Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", CPU.Hz, "hz")

	// Test if we have these specific features:
	if CPU.Supports(SSE, SSE2) {
		fmt.Println("We have Streaming SIMD 2 Extensions")
	}
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
