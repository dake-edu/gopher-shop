# Professional Tools: The Auto-Pilot

Professional developers don't check everything manually. They use tools to do it for them.

## The "Green Build"

When you save your game, you want to know if you made progress or if you broke something.
In coding, we have a **Continuous Integration (CI)** pipeline.

### What is it?
Ideally, every time you send your code to GitHub, a robot wakes up and:
1. **Tries to Build it**: Does the code even compile?
2. **Tests it**: Do all the tests passed?

If everything is good, the robot gives you a **Green Checkmark** ✅. This is a "Green Build".
If something is wrong, you get a **Red X** ❌.

## Our Setup

We have configured a robot (GitHub Actions) in `.github/workflows/go-check.yml`.
It runs:
- `go build ./...` (Structure Check)
- `go test ./...` (Functionality Check)

### Why is this important?
It gives you **CONFIDENCE**. You can sleep effectively knowing that the robot checked your work.
