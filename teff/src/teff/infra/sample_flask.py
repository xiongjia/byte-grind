#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Filename: sample_flask.py
"""

from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello():
   return 'hello'

def test_main():
  app.run("127.0.0.1", 8083)

if __name__ == "__main__":
  test_main()
