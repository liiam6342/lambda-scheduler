## lambda-scheduler
This is the lambda handler code

**getInstances.go**
Looks for instances using the tag "Scheduler" with a value of "True" 

**startInstances.go**
Starts the specific instance if it is not already running

**stopInstances.go** 
Stops the specific instance if it is running

**scheduler.go**
This file checks the current time and decides the action
 
