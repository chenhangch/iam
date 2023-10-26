# Commit Message 规范

Go开发项目中，一个清晰的、好的Commit Message规范有以下作用：
+ 帮助开发人员清晰地知道每个提交 （commit）的变更内容，方便浏览变更历史
+ 基于Commit Message进行查找过滤，例如：查找某个版本新增功能
    ```cmd
  git log --oneline --grep "^feat|^fix|^perf"
  ```
+ 可以基于规范化的Commit Message生成变更日志
+ 可以依据某些类型Commit Message触发构建或者发布流程
+ 可以用来确定语义化版本的版本号

## 符合Angular规范的Commit Message
+ Commit Message是语义化的：Commit Message都会被归于一个有意义的类型，说明本次提交的类型
+ Commit Message是规范化的：Commit Message遵循预先定义好的规范

在Angular规范下，Commit Message包含三个部分，分别是**Header**、**Body**和**Footer**，格式如下：
```cmd
<type>(<scope>): <subject>
// 空行
[optional body]
// 空行
[optional footer(r)]
```
`type`：分为两种类别: **Development** && **Production**
+ Deployment：这类修改一般是项目管理类的变更，不会影响到用户和生产环境的代码
+ Production：这类修改会影响到生产环境的代码，对于这种类型的改动，提交时需要慎重，并在提交前做好充分的测试

常见类型以及它们所属的类别：

| 类型       | 类别          | 说明                                                  |
|----------|-------------|-----------------------------------------------------|
| feat     | Production  | 新增功能                                                |
| fix      | Production  | 修复缺陷                                                |
| perf     | Production  | 提高代码性能的变更                                           |
| style    | Development | 代码格式类的变更，例如使用gofmt格式化代码                             |
| refactor | Production  | 其他代码类的变更，例如 简化代码、重命名变量、删除冗余代码等等                     |
| test     | Development | 新增测试用例或更新现有的测试用例                                    |
| ci       | Development | 持续基础和部署相关的改动，例如修改Jenkins、GitLab CI等Ci配置文件或者更新系统单元文件 |
| docs     | Development | 文档类的更新，包括修改用户文档、开发文档                                |
| chore    | Development | 其他类型，例如构建流程、依赖管理或者复制工具的变动                           |

`scope`：用以说明commit的影响范围

`subject`：是commit的简短描述，必须以动词开头，使用现在时。明确地描述本次commit的内容

`Body`：对commit做高度概况，是详细的，格式相对比较自由的描述

`Footer`：非必选，主要用来说明本次commit导致的后果。**不兼容的改动**或者**关闭Issue列表**

特殊的`Revert Commit`：如果当前的commit还原了先前的commit，则应以```revert:```开头，后面跟还原的commit的Header
Body必须写上`This reverts commit <hash>`，其中`hash`是还原的commit的**SHA**标识

## Commit Message的三个重要内容
+ 提交频率
  1. 对项目进行了修改，通过测试便可以提交
  2. 定时提交，例如每天结束写代码时进行一次提交
  3. 在上面的提交中可能出现次数过多的情况， 在最后的合并代码或者提交Pull Request前，执行`git rebase -i`合并之前所有commit
+ 合并提交：将多个commit合并成为一个commit提交
  1. 使用`git rebase`命令进行合并
+ Commit Message修改
  1. `git commit --amend`：修改最近一次的Commit Message
  2. `git rebase --i`：修改某次Commit Message

## Commit Message规范自动化
自动化工具

| 工具名           | 功能                                             |
|---------------|------------------------------------------------|
| commitizen-go | 进入交互模式，并根据提示生成Commit Message，然后提交              |
| commit-msg    | githooks，在commit-msg中知道检查的规则。是一个脚本，可以根据需要写脚本实现 |
| go-gitlint    | 检查历史提交是否符合Angular规范                            |
| gsemver       | 语义化版本自动生成工具                                    |
| git-chglog    | 根据Commit Message自动生成CHANGELOG                  |

### 版本规范
+ 主版本号（MAJOR）：在了不兼容的API修改时递增 -->[BREAKING CHANGE]
+ 此版本号（MINOR）：做了向下兼容的功能性新增以及修改时递增 -->[feat]
+ 修订号（PATCH）：做了向下兼容的问题修改时递增 -->[fix]
