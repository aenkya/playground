from django.shortcuts import render
from django.contrib.auth.models import User

def signup(request):
    if request.method == 'POST':
        pass
    return render(request, 'users/signup.html')