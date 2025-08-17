#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Filename: util.py
Description: test pytest 
"""

import pytest
from teff.infra.util import add_one

def test_add_one():
    result = add_one(3)
    assert result == 4, "!= 4"

def test_two():
    assert 1 == 1, "!= 1"
