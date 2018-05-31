import logging

#from django.http import HttpResponse  
from django.http.response import JsonResponse
from .models import ResultTO
  
logger = logging.getLogger(__name__)

# http api  
def getResult(request, id):
    logger.debug('Enter')
    result = ResultTO()
    result.return_code = id
    result.message = 'hello'
    result.success = True    
    return JsonResponse(result.__dict__)
