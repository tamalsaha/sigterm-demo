# sigterm-demo

https://en.wikipedia.org/wiki/Process_group

In a POSIX-conformant operating system, a process group denotes a collection of one or more processes.[1] Among other things, a process group is used to control the distribution of a signal; when a signal is directed to a process group, the signal is delivered to each process that is a member of the group.[2]

Similarly, a session denotes a collection of one or more process groups.[3] A process may not create a process group that belongs to another session; furthermore, a process is not permitted to join a process group that is a member of another session—that is, a process is not permitted to migrate from one session to another.

When a process replaces its image with a new image (by calling one of the exec functions), the new image is subjected to the same process group (and thus session) membership as the old image.

The system call setsid is used to create a new session containing a single (new) process group, with the current process as both the session leader and the process group leader of that single process group.[4] Process groups are identified by a positive integer, the process group ID, which is the process identifier of the process that is (or was) the process group leader. Process groups need not necessarily have leaders, although they always begin with one. Sessions are identified by the process group ID of the session leader. POSIX prohibits the change of the process group ID of a session leader.

The system call setpgid is used to set the process group ID of a process, thereby either joining the process to an existing process group, or creating a new process group within the session of the process with the process becoming the process group leader of the newly created group.[5] POSIX prohibits the re-use of a process ID where a process group with that identifier still exists (i. e. where the leader of a process group has exited, but other processes in the group still exist). It thereby guarantees that processes may not accidentally become process group leaders.

The system call kill is capable of directing signals either to individual processes or to process groups.[2]
