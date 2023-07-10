## A Go Cron Scheduler for Humans.

Basically a scheduler which takes a queue of jobs and executes them in batches. The batch size mentioned here is 3 so at a time only 3 jobs will be executed. The scheduler will keep on executing the jobs in batches until the queue is empty. There is a main.py file for testing as well.