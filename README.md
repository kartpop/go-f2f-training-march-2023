# Dependancy Management

# Logging

# Configuration
    - Webhost

# Debug and Metrics
http://localhost:4000/debug/vars
http://localhost:4000/debug/pprof

# shutdown signaling
    -



 #WebFramework
 - Handler customization


 Problem/Task
 New signature support needs to be added

 // type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
 
// func (c check) readiness(w http.ResponseWriter, r *http.Request) error {
 API argument/signature we can't change to pass logger so use struct and attached to method
 Also I have to return error via handler

 How ?
 1. Add web package in handler
 2. Add new type Handler signature in web package
    type Handler func(ctx context.Context, shutdown chan os.Signal, log *log.Logger) error


Task pass shutdown to app
//shutdown                chan os.Signal



// Tracing info
 add context info in web.go
 replace r.context() 
 	// Set the context with the required values to
		// process the request.
		v := Values{
			TraceID: span.SpanContext().TraceID.String(),
			Now:     time.Now(),
		}
		ctx = context.WithValue(ctx, KeyValues, &v)

        https://github.com/google/uuid