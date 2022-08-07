package main

type Server struct {
	recordTime string
	ip string
	responseTime string
	failuare bool
}

type BreakServer struct {
	ip string
	breakStartTime string
	breakEndTime string
}