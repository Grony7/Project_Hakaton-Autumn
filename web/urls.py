from . import views
from django.contrib import admin
from django.contrib.staticfiles.urls import staticfiles_urlpatterns
from django.urls import path, include
from django.views.generic import RedirectView

urlpatterns = [
  path('', views.index, name='general'),
  path('login/', views.login, name='login')

]


urlpatterns += staticfiles_urlpatterns()
