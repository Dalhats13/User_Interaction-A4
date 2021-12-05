module main

require Posting-Feed-be/transport v0.0.0

require Posting-Feed-be/datastruct v0.0.0 //indirect

require Posting-Feed-be/logging v0.0.0 //indirect

require (
	Posting-Feed-be/service v0.0.0 //indirect
	github.com/lib/pq v1.10.3
)

replace Posting-Feed-be/transport => ./transport

replace Posting-Feed-be/datastruct => ./datastruct

replace Posting-Feed-be/logging => ./logging

replace Posting-Feed-be/service => ./service

require (
	github.com/go-kit/kit v0.12.0
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
)

go 1.17
