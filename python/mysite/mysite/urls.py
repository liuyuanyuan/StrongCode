"""mysite URL Configuration

The `urlpatterns` list routes URLs to http_views. For more information please see:
    https://docs.djangoproject.com/en/2.0/topics/http/urls/
Examples:
Function http_views
    1. Add an import:  from my_app import http_views
    2. Add a URL to urlpatterns:  path('', http_views.home, name='home')
Class-based http_views
    1. Add an import:  from other_app.http_views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
#from django.contrib import admin
#from django.urls import path
#from django.conf.urls import include,
from django.conf.urls import url
from mysite import http_views
from mysite import http_api
 
urlpatterns = [
    # path('admin/', admin.site.urls),
    #url(r'^admin/', include(admin.site.urls)),
    url(r'^$', http_views.output), #browse http://192.168.100.172:8000/
    url(r'^getResult/(?P<id>[0-9]+)/$', http_api.getResult, name='getResult') # http://192.168.100.172:8000/getResult/1/
]




