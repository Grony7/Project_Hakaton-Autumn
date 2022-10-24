from django import forms
from django.forms import ModelForm

from web.models import User


class LoginForm(ModelForm):
  class Meta:
    model = User
    fields = ['name', 'email', 'password', 'who']
