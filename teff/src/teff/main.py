#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Filename: main.py
Description: test pytest
"""

from teff.infra.util import add_one

def main():
  print("test %d\n" % add_one(100))

if __name__ == "__main__":
  main()

