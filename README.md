# rldemo

## Overview

This is a demo project with a particular target audience; you know who you are!

## Getting Started

> Skip the setup and view it live at: [https://rldemo.santee.cloud](https://rldemo.santee.cloud)

Copy or rename `.env.example` to `.env.local`.

Run:

``` bash
pnpm install
# then
pnpm dev
```

### Local API testing

In the env.local, swap out the `santee.cloud` endpoints for the `localhost` endpoints, then in another console window, run 
```bash
npm run api
```

> Don't have `pnpm`? [Find it here](https://pnpm.io/installation).

> Don't want `pnpm`? `npm install` and `npm run dev` will have the same effect!

You should now be able to access the application at <http://localhost:3000>.

## Architecture

The application consists of a NextJS front-end/back-end, and serverless services in Go implementing the core business logic. It is designed to deploy to [Vercel](https://vercel.com/home).

The project started with the [NextJS Admin Dashboard](https://vercel.com/templates/next.js/admin-dashboard) template (you may see a few traces of the boilerplate hanging around).

Illustration of the project structure:
![Project Structure](./docs/project-structure.png "Project Stucture")

## Requirements

### Where are they implemented?

They are split between the UI and the serverless services.

The services should be fairly easy to navigate and understand. There are two, they live under `/api`.

`/api/fbz` implements the FizzBuzz logic, and `/api/fib` implements the Fibonacci logic.

The UI is implemented in the NextJS app, in TypeScript, with React, supported by assorted libraries. The pages live under `/app/(dashboard)`.

Be sure to look for the gear icon in each UI, as some requirements are hiding there!

| Requirement | Implemented at             | View Live                               |
|-------------|----------------------------|-----------------------------------------|
| Part 1      | /app/(dashboard)/fizzbuzz  | <https://rldemo.santee.cloud/fizzbuzz>  |
| Part 2      | /app/(dashboard)/fibonacci | <https://rldemo.santee.cloud/fibonacci> |
| Part 3      | /app/(dashboard)/fizzbuzz  | <https://rldemo.santee.cloud/fizzbuzz>  |

### Why this stack?

The stack was chosen because it is my understanding that you are interested in NextJS. It had been a while since I worked with React/front-end frameworks, so I though it would be a good opportunity to refresh my own knowledge, and hopefully illustrate competency.

Vercel was new to me. I chose it because it promised ease, which I thought would be an important factor in delivering this project in a reasonable time frame.

### Why the services?

It's often beneficial to split core business logic out into separate layers that are indepently deplpoyable and reusable. I thought this design would illustrate some architectural know-how and concern.

Use of Go was both for illustration of flexibility, and expedience, as I've worked with it recently and could move relatively quickly.

## Notes

Ran into this issue with apparent misconfiguration of tsconfig.json.
<https://github.com/shadcn-ui/ui/issues/1092>

## TODO
 - (Dashboard) route group is unnecessary
   - https://nextjs.org/docs/app/getting-started/project-structure#routing-files
 - Remove a `types.ts` from one of `/fibonacci` or `/fizzbuzz` and use a common copy
 - customize `/app/error.tsx/`
 - extract `about.tsx` to a common component that takes props
    - https://react.dev/learn/passing-props-to-a-component#step-2-read-props-inside-the-child-component
 - UI should reflect `1000000` maximium input to fibonacci