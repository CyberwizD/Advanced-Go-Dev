package backgroundjob

// Background Job is a process in charge of doing some work "behind the scenes"
// Initialized by another "parent" process, which simply means that a goroutine launching `N` goroutines
