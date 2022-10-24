from django.contrib import admin
from web.models import User, Feedback, News

# Register your models here.
admin.site.register(User)
admin.site.register(Feedback)
admin.site.register(News)
