# harbor-mob-offline-lens

`harbor-mob-offline-lens` explores mobile workflows with a small Go codebase and local fixtures. The technical goal is to create a Go reference implementation for offline workflows, centered on stream reduction, windowed input fixtures, and late-data behavior checks.

## Problem It Tries To Make Smaller

The project exists to keep a narrow engineering decision visible and testable. For this repo, that decision is how form pressure and local state should influence a review result.

## Harbor Mob Offline Lens Review Notes

The first comparison I would make is `local state` against `form pressure` because it shows where the rule is most opinionated.

## Working Pieces

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/harbor-mob-offline-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `local state` and `form pressure`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Design Notes

The repository has two validation layers: the original compact policy fixture and the domain review fixture. They are separate so one can change without hiding failures in the other.

The added Go path is deliberately direct, with fixtures doing most of the explaining.

## Example Run

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Tests

The verifier is intentionally local. It should fail if the fixture score math, lane assignment, or language-specific test drifts.

## Known Limits

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
