#!/bin/bash

SERVICE_NAME="wifi_plugins"
START_COMMAND="/etc/init.d/wifi_plugins_init start"

# 指定日志文件路径
LOG_FILE="wifi_plugins.log"

# 检查日志文件是否存在，不存在则创建
[ ! -f "$LOG_FILE" ] && touch "$LOG_FILE"

# 持续循环检查服务状态
while true; do
    # 使用ps命令或其他方式检查服务是否正在运行
    if ! pgrep -x "./wifi_plugins" > /dev/null; then
        echo "$(date) - 服务[$SERVICE_NAME]未检测到运行状态，将尝试重新启动..." >> $LOG_FILE

        # 尝试启动服务
        $START_COMMAND

        # 这里可能需要添加一些逻辑来确认服务是否成功启动
        # 对于SysVinit脚本，可以通过检查启动后的输出或者等待一小段时间后再次检查服务进程是否存在
        sleep 180
        if ! pgrep -x "./wifi_plugins" > /dev/null; then
            echo "$(date) - 服务[$SERVICE_NAME]重新启动失败，请检查相关日志。" >> $LOG_FILE
        else
            echo "$(date) - 服务[$SERVICE_NAME]已成功重新启动。" >> $LOG_FILE
        fi
    fi
    # 延迟一段时间后再次检查（例如每30秒检查一次）
    sleep 900
done