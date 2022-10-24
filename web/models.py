from django.db import models

# Create your models here.
from django.urls import reverse


class Idk(models.Model):
  name = models.CharField


def add_user(name, email, password, who):
  record = User(name=name, email=email, password=password, who=who)
  record.save()


class User(models.Model):
  id = models.AutoField(primary_key=True)
  name = models.CharField(max_length=150)
  email = models.EmailField(max_length=30)
  password = models.CharField(max_length=40)
  who = models.CharField(max_length=20)

  def __str__(self):
    return self.name

  def get_email(self):
    return self.email

  def get_password(self):
    return self.password

  def get_absolute_url(self):
    return reverse('user', args=[str(self.id)])


class Feedback(models.Model):
  name = models.CharField(max_length=70)
  email = models.EmailField(max_length=100, null=True)
  text = models.CharField(max_length=1000)

  def __str__(self):
    return self.email


def add_feedback(name, email, text):
  record = Feedback(name=name, email=email, text=text)
  record.save()


class News(models.Model):
  title = models.CharField(max_length=200)
  main_text = models.CharField(max_length=10000)
  preview = models.CharField(max_length=150)
  public_date = models.DateField(null=True)
  image = models.ImageField(upload_to='files/', blank=True)

  def __str__(self):
    return self.title


def upload_news(title, main_text, preview, public_date, image):
  record = News(title, main_text, preview, public_date, image)
  record.save()
