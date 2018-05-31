# -*- coding: utf-8 -*-
from __future__ import unicode_literals
#from django.db import models

# Create data model here
class ResultTO:
    def __init__(self):
        self.return_code = None
        self.success = False
        self.message = None
        self.data = None
        self.data_list = None
        
        