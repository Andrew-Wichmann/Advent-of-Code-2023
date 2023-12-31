# Advent-of-Code-2023

## Day 1

This took me a couple of hours. I'm trying out GitHub CodeSpaces and GitHub CodePilot, so that played into it. I'm also using GoLang which I'm not the most familiar with.

I think that doing it in python would make this a lot less verbose. Also, I found myself casting chars into strings and runes a lot. Casting into runes was probably overkill since we're not using unicode codepoints anyway. It might be neat to try the same challenge with some unicode codepoints (is that the same as a "character"?) thrown in.

I can think of one optimization; in `mapStringsToNumbers` the second loop iterates through the whole rest of the string, but, really, it only needs to check the next 4 characters because there are no numbers that, when spelled out, are longer than 5 characters. Ex: given the string "abcdefg", there's no need to check beyond "abcde" because "eight" is the longest word from 0-9.

I have a feeling that having to create a test for `mapStringsToNumbers` means I did something wrong. Like, this change should be simple enough that the code should need tests to explain it.

## Day 2

EZPZ. Very fun though. This is truly an amazing way to learn a new language. These feel like perfectly approachable problems, and you get great practice at reading files, parsing strings, creating objects, and writing algorithms.

I might want to create like a cookiecutter or something to do the boilerplate.

I'm LOVING codespaces. I love love love it, and I really wish we could have something like this in GitLab. Oh [there is](https://docs.gitlab.com/ee/user/workspace/)? I'll have to see if we have this feature.
