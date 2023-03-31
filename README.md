# Parsing C Kernel Source Files with Deterministic Grammar

This repository contains a demonstration of the limitations of parsing C 
kernel source files using a deterministic grammar. While it is possible 
to parse C and C++ languages separately with a deterministic grammar, 
the combination of C and C preprocessor (cpp) can make parsing together 
very difficult.

The main reason for this is not the complexity of the languages, but 
the use of the C preprocessor, especially ifdefs and macros. When C 
and cpp are combined, they create a union of two languages that are 
easy to parse separately but very hard to parse together. 

The resulting grammar could be very large, contain many ambiguities, 
and be very different from the original C grammar.

While tools like Coccinelle semantic patch use a lot of heuristics 
to overcome these problems, the effort of implementing the same 
compared with the benefits discourages this practice.

This repository serves as a proof of concept for the limitations 
of parsing C kernel source files with a deterministic grammar. 
It demonstrates the challenges that arise when attempting to parse 
a language that includes both C and cpp, and highlights the need 
for alternative approaches to parsing C kernel source files.

**The demonstrator is able to parse not simple  preprocessed c 
files but it is unable to cope with complex macros.**

