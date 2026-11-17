# Week 1 evidence

> **How to use this file.** Fill in every `[ ... ]` with output from **your own machine**. Do not copy expected output from the guide: the reviewer needs to see what happened on your computer. Delete this note before you share the file.

| | |
| --- | --- |
| Name | [your name] |
| Date | [date] |
| Operating system | [for example: Windows 11] |
| `go version` | [paste the line] |
| `temporal --version` | [paste the line] |

The onboarding guide (page 3) asks for this evidence in Week 1:

> Explain page 2; run a workflow and activity locally.

---

## Part A — Explain page 2

I can explain every term in the table on page 2 without notes. My answers to the three completion-check questions, in my own words:

**1. Why does a successful token request not prove access to every API?**

[your answer, 2–4 sentences]

**2. Why is a group object not its complete membership list?**

[your answer, 2–4 sentences]

**3. Why does missing policy data not mean that no policy exists?**

[your answer, 2–4 sentences]

---

## Part B — A workflow and an activity, run locally

**What I ran.** A Temporal development server, one worker and one starter on my own machine. The workflow (`CollectDemoWorkflow`) calls one activity (`LoadSamplePage`), waits on a timer, and returns a small result. The activity is written to fail on its first two attempts, so the run also shows Temporal retrying it.

Commands, each in its own terminal window:

```
temporal server start-dev
go run ./lesson3/worker
go run ./lesson3/starter
```

**Starter output:**

```
[paste the starter output here]
```

**Worker output (the two failed attempts):**

```
[paste the two "simulated network error" lines here]
```

**Workflow history** (`temporal workflow show --workflow-id week1-demo`):

```
[paste the event list here]
```

**Screenshot of the Temporal Web UI** (`http://localhost:8233`, workflow `week1-demo`): [attach the image]

**What this shows, in my own words:**

- Which line of the history is the activity: [your answer]
- Why the activity appears once in the history although it was attempted three times: [your answer]
- What the difference between a workflow and an activity is: [your answer]

---

## Part C — Worker restart (optional, recommended)

I started the workflow with a 60-second timer, stopped the worker while the timer was running, waited, and started the worker again.

```
[paste the starter output here]
```

- Did the workflow finish after the worker came back? [yes / no]
- Did `LoadSamplePage` run again after the restart? [yes / no — how you can tell]
- Why: [your answer]

---

## Part D — Go: JSON, HTTP, context and tests

Output of `go test ./... -v` filtered to my exercises:

```
[paste the lines for TestExercise0, TestExercise1, TestExercise2 and TestExercise3]
```

---

## Notes for the reviewer

- All data in these lessons is synthetic. No real tenant, credential or token was used in Week 1.
- The lesson code was prepared with an AI assistant as learning material. The exercise solutions are my own work: [yes / partly — say which].
- Not done yet: any call to a real Entra tenant. I need the development tenant and repository access to start Week 2.