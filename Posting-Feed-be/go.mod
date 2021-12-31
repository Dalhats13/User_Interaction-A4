module main

require Posting-Feed-be/transport v0.0.0

require Posting-Feed-be/datastruct v0.0.0 //indirect

require (
	Posting-Feed-be/logging v0.0.0 //indirect
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/gofiber/fiber/v2 v2.23.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.31.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678 // indirect
)

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
