from django.shortcuts import render
from .models import User, add_user
from .forms import LoginForm


# Create your views here.

def index(request):
  name = request.GET.get('feedback-name')
  email = request.GET.get('feedback-mail')
  password = request.GET.get('feedback-pass')
  who = request.GET.get('who')
  if request.method == 'GET' and name is not None:
    form = LoginForm()

    add_user(name, email, password, who)
  return render(request, 'index.html')


def login(request):
  return render(request, 'register.html')
