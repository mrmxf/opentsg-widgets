# pipe in the $ENV for bot-tlh - just runs with enviroment variables
aws ec2 start-instances --instance-ids i-0dc45d0982ea473f8 

# give bot tlh the ability to stop start instances
aws ec2 stop-instances --instance-ids i-0dc45d0982ea473f8 