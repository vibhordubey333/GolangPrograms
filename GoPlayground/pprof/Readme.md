##### What is memory leak ?

Memory leaks are a class of bugs where memory is not released even after it is no longer needed. They are often explicit, and highly visible, which makes them a great candidate to begin learning debugging. Go is a language particularly well suited to identifying memory leaks because of its powerful toolchain, which ships with amazingly capable tools (pprof) which make pinpointing memory usage easy.
