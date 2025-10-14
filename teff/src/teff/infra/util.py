#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Filename: util.py
"""

import functools

def add_one(number):
  return number + 1

def simple_decorator(func):
  @functools.wraps(func)
  def wrapper(*args, **kwargs):
    print(f"Calling function: {func.__name__}")
    result = func(*args, **kwargs)
    print("After the function call. %d" % result)
    return result
  return wrapper

def simple_decorator2(func):
  @functools.wraps(func)
  def wrapper(*args, **kwargs):
    print(f"Calling function: {func.__name__}")
    result = func(*args, **kwargs)
    print("After the function call. %d" % result)
    return result
  return wrapper

@simple_decorator2
@simple_decorator
def greet(a, b = 1):
  """ greet """
  print("hello")
  return a + b

def test_main():
  result = greet(3)
  print("result %d" % result)

  print("greet func and doc:")
  print(greet.__name__)
  print(greet.__doc__)

if __name__ == "__main__":
  test_main()

