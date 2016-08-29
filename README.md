Barber
======

Have you ever wanted to run a command and insert the output without a
newline?

Barber trims the new lines of the command output.


Usage
=====

When you want to strip the newlines at the end of a command, just
prefix it with `barber`.

```bash
$ barber date +%Y-%m-%d
```

That's it.


Why?
====

There a tons tools out there that let you put the result from a script
or command somewhere helpful. For example, when I was able to use
linux, I had a few little scripts tell me some system information in
my [StumpWM](https://stumpwm.github.io/) menu bar. Similarly, I use
[Emacs](https://www.gnu.org/software/emacs/) and have written small
elisp functions to insert the result of commands. When using this
pattern, I would often throw up my hands in frustration because after
putting together some command, I'd have a newline in the output,
ruining the result.

So, I wrote this little tool to fix it. I can imagine it could spawn
other features such as allowing to you mimic piping where you don't
have an actual shell, but I'll cross that bridge when I reach it. Hope
it helps someone else!
