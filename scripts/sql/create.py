import os

import mysql.connector
import yaml

# 加载yaml配置文件中的数据库信息
with open("db.yaml", "r") as file:
    db_config = yaml.safe_load(file)

# 连接数据库
mydb = mysql.connector.connect(
    host=db_config['mysql']['host'],
    port=db_config['mysql']['port'],
    user=db_config['mysql']['user'],
    passwd=db_config['mysql']['passwd'],
    database=db_config['mysql']['database']
)

mycursor = mydb.cursor()

directory = 'tables'

# 遍历表格
for table in db_config['tables']:
    table_name = table['name']
    count = table['count']

    # 打开并遍历SQL文件
    with open(os.path.join(directory, f'{table_name}.sql'), 'r') as file:
        sql = file.read()

        for i in range(1, int(count) + 1):
            if count == 1:
                real_table_name = table_name
            else:
                real_table_name = table_name + '_' + str(i)

            # 将sql语句中的表名替换为真实的表名
            real_sql = sql.replace('CREATE TABLE `' + table_name + '`', 'CREATE TABLE `' + real_table_name + '`')

            # 执行sql语句
            mycursor.execute(real_sql)

mycursor.close()
mydb.close()
