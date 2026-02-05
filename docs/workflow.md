# 16. The Robot (CI/CD)

The most professional developers don't deploy manually. They have a robot do it.

## 16.1 Continuous Integration (CI)
**"The Gatekeeper"**
Every time you push code to GitHub, a robot (GitHub Actions) wakes up.
1. It downloads your code.
2. It tries to build it (`go build`).
3. It runs all tests (`go test`).

If *anything* fails, the robot blocks the Merge. **Green Build** = Safe Code.

## 16.2 Continuous Deployment (CD)
**"The Delivery Truck"**
If the CI is Green, the CD robot takes the artifact (the binary or documentation) and ships it to the world (Production Server or GitHub Pages).

We are using CD right now to publish this documentation site!
