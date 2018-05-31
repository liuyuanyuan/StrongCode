#!/usr/bin/env python
import os
import sys

if __name__ == "__main__":
    os.environ.setdefault("DJANGO_SETTINGS_MODULE", "mysite.settings")
    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError(
            "Couldn't import Django. Are you sure it's installed and "
            "available on your PYTHONPATH environment variable? Did you "
            "forget to activate a virtual environment?"
        ) from exc
    execute_from_command_line(sys.argv)
    #execute_from_command_line(['manage.py','runserver','0.0.0.0:8000'])
    
'''
Start server as following:
E:\EclipseWorkspace\mysite>python manage.py runserver 0.0.0.0:8000
Performing system checks...

System check identified no issues (0 silenced).
May 31, 2018 - 13:37:18
Django version 2.0.5, using settings 'mysite.settings'
Starting development server at http://0.0.0.0:8000/
Quit the server with CTRL-BREAK.
'''
