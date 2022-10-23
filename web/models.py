from django.db import models

# Create your models here.
from django.urls import reverse


class User(models.Model):
  id = models.AutoField(primary_key=True)
  name = models.CharField(max_length=150)
  email = models.EmailField
  password = models.CharField
  who = models.IntegerField

  def __str__(self):
    return self.name

  def get_email(self):
    return self.email

  def get_password(self):
    return self.password

  def get_absolute_url(self):
    return reverse('user', args=[str(self.id)])


record = User(name='bruh')
record.save()


class Meta:
  ordering = ['id']
