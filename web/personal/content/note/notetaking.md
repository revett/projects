---
title: "Notetaking"
draft: true
---

## Leitner System

A flashcard system based on spaced repetition, where cards are reviewed at
different intervals. First outlined by Sebastian Leitner in his 1972 book titled
"How to Learn to Learn".

Methodology:

- There are `n` boxes, each representing a different proficiency level
- At a given point, a specific card will be in one of the `n` boxes
- All cards start in the first box (e.g. `A`)
- When the user successfully recalls the solution to the card, the card
  is moved to the next box (e.g. `A → B`)
- If they fail to recall the solution, the card is moved back to the first
  box
- The box sets how often a card is reviewed, the further away from the starting
  box the more infrequently the card is studied

Note that an alternative method sees incorrect cards only moved back one box,
instead of being sent all the way back to the first box.
