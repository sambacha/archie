---
title: "Sweet Pandas Eating Sweets"
date: "2017-08-10"
---

# Archie documentation

## Overview

Archie is a lightweight tool for generating model-based system architecture diagrams.

It's features:

- All diagrams generated from a yaml model
  - Actor elements (users)
  - Item elements (modules) that may be infinitly composed (sub-elements)
  - Associations between any elements
  - Elements and associations may be tagged
- Various diagrams
  - Landscape view for reviewing root elements
  - Context view for reviewing an item and it's interfaces
  - Tag view for reviewing a type of item in the system
- Various methods for invocation
  - A CLI for local generation
  - A REST API for remote generation/interfacing

## Installation

Install the archie CLI with:

```bash
go get github.com/briggysmalls/archie/cli/archie
```

## CLI Documentation

Then read the [documentation](cli/archie.md) for invoking it!