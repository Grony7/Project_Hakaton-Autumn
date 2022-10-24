from django.core.mail import send_mail
from django.shortcuts import render
from .models import User, add_user, add_feedback


# Create your views here.

def index(request):
  name = request.POST.get('feedback-name')
  email = request.POST.get('feedback-mail')
  message = request.POST.get('feedback-message')
  if request.method == 'POST' and name is not None:
    add_feedback(name, email, message)
    send_mail(subject='Ректору', message=message,
              from_email='rector.site@gmail.com',
              recipient_list=[email])
  return render(request, 'index.html')


def register(request):
  name = request.POST.get('field-name')
  email = request.POST.get('field-email')
  error = 0
  try:
    user = User.objects.get(email=email)
    error = 1
  except User.DoesNotExist:
    user = None
  password = request.POST.get('field-pass')
  who = request.POST.get('product-group')
  if request.method == 'POST' and name is not None and user is None:
    add_user(name, email, password, who)
  return render(request, 'register.html', context={'error': error})


def login(request):
  email_form = request.POST.get('field-email')
  user = None
  if email_form is not None:
    try:
      user = User.objects.get(email=email_form)
    except User.DoesNotExist:
      user = None
  result = 0
  if user is not None:

    password_form = request.POST.get('field-pass')
    if password_form == user.password:
      result = 1
    else:
      result = 2
  context = {
    'result': result
  }
  if result == 1:
    return render(request, 'index.html', context)
  else:
    return render(request, 'login.html', context)
