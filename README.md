# sigterm-demo

- https://github.com/tamalsaha/learn-bash/issues/31
- https://github.com/appscode/discuss/issues/393


https://en.wikipedia.org/wiki/Process_group

In a POSIX-conformant operating system, **a process group denotes a collection of one or more processes.** Among other things, a process group is used to control the distribution of a signal; **when a signal is directed to a process group, the signal is delivered to each process that is a member of the group.**

Similarly, a session denotes a collection of one or more process groups. A process may not create a process group that belongs to another session; furthermore, a process is not permitted to join a process group that is a member of another session—that is, a process is not permitted to migrate from one session to another.

**When a process replaces its image with a new image (by calling one of the exec functions), the new image is subjected to the same process group (and thus session) membership as the old image.**

The system call setsid is used to create a new session containing a single (new) process group, with the current process as both the session leader and the process group leader of that single process group. Process groups are identified by a positive integer, the process group ID, which is the process identifier of the process that is (or was) the process group leader. Process groups need not necessarily have leaders, although they always begin with one. Sessions are identified by the process group ID of the session leader. POSIX prohibits the change of the process group ID of a session leader.

The system call setpgid is used to set the process group ID of a process, thereby either joining the process to an existing process group, or creating a new process group within the session of the process with the process becoming the process group leader of the newly created group. POSIX prohibits the re-use of a process ID where a process group with that identifier still exists (i. e. where the leader of a process group has exited, but other processes in the group still exist). It thereby guarantees that processes may not accidentally become process group leaders.

**The system call kill is capable of directing signals either to individual processes or to process groups.**


## Foreground process from go binary -> script.sh -> test program (sleep)

```
2899232 2745605 2899232 2899232 bash
2918303 2899232 2918303 2899232 sigterm-demo
2918308 2918303 2918303 2899232 script.sh
2918309 2918308 2918303 2899232 sleep
```

>> Bash does not pass SIGTERM

## Background process (SIGINT or Ctrl+C)

ref: https://unix.stackexchange.com/a/362566

```
$ ps  xao pid,ppid,pgid,sid,comm | grep 2899232
2899232 2745605 2899232 2899232 bash
2918470 2899232 2918470 2899232 sigterm-demo
2918475 2918470 2918470 2899232 script.sh
2918476 2918475 2918470 2899232 sleep

# Hit ctrl+C in keyword (SIGINT)

$ ps  xao pid,ppid,pgid,sid,comm | grep 2899232
2899232 2745605 2899232 2899232 bash
2918476       1 2918470 2899232 sleep
```

## Background process (SIGTERM)

ref: https://www.linuxcertified.com/e-learning/linuxplus/html/5_7.html

```
$ ps  xao pid,ppid,pgid,sid,comm | grep 2899232
2899232 2745605 2899232 2899232 bash
2919145 2899232 2919145 2899232 sigterm-demo
2919150 2919145 2919145 2899232 script.sh
2919151 2919150 2919145 2899232 sleep

# SIGTERM to sigterm-demo
kill -15 2919145

$ ps  xao pid,ppid,pgid,sid,comm | grep 2899232
2899232 2745605 2899232 2899232 bash
2919150       1 2919145 2899232 script.sh
2919151 2919150 2919145 2899232 sleep

# SIGTERM to script.sh
$ kill -15 2919150

$ ps  xao pid,ppid,pgid,sid,comm | grep 2899232
2899232 2745605 2899232 2899232 bash
```

>> Looks like SIGTERM not getting forwarded from `sigterm-demo` to `script.sh`

>> Should `sigterm-demo` block for child processes to exit
