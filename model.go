package main

type Server struct {
	recordTime string
	ip string
	responseTime string
	failuare bool
	timeoutCount int
}

type BreakServer struct {
	ip string
	breakStartTime string
	breakEndTime string
}

type TimeoutServer struct {
	ip string
}

type SubnetServer struct {
	subnet string
}

type BreakSubnetServer struct {
	subnet string
	breakStartTime string
	breakEndTime string
}