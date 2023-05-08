> 该logger直接从marmotedu/iam/internal/pkg/logger进行重构

该logger对gorm.logger进行重写，定义了也该Config的结构体，用于配置logger的行为

然后，实现了gormlogger.Interface接口，该接口定义了四个方法：LogMode、Info、Warn、Error
用以记录不同级别的日志信息
