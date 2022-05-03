# `/backend/mainmodule/internal`

Protected package code for the app.

If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!

*NOTE: The Go compiler enforces that no package in an `internal/` directory can be imported by any other application or pacakge that does nto live in the same parent path is the internal directory.*
