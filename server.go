package main

type Server struct {
	recordTime string
	ip string
	responseTime string
	breakTime int
	timeoutCount int
}

type ServerResponse struct {
	ip string
	responseTime string
}