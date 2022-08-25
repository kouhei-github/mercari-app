FROM mysql:5.7

COPY my.cnf /etc/mysql/conf.d/my.cnf

# 日本語環境を追加
RUN apt-get update && apt-get install -y locales \
    && sed -i -e 's/# \(ja_JP.UTF-8\)/\1/' /etc/locale.gen \
    && locale-gen \
    && update-locale LANG=ja_JP.UTF-8
# ここまで

ENV LC_ALL ja_JP.UTF-8

ENV TZ Asia/Tokyo
ENV LANG=ja_JP.UTF-8

CMD ["mysqld"]

EXPOSE 3306
