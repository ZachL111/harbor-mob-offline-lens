# Harbor Mob Offline Lens Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 107 | watch |
| stress | sync drift | 161 | ship |
| edge | local state | 172 | ship |
| recovery | conflict cost | 151 | ship |
| stale | form pressure | 167 | ship |

Start with `edge` and `baseline`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

If `baseline` becomes less cautious without a clear reason, I would inspect the drag input first.
