# ExcelBuilder
java语言有EasyPoi可用于简单、快速的生成excel文件，golang语言
目前只有各种excelize用于excel文件的生成，但有一个坏处就是不提供
类似于esayPoi高度集成、封装的方法，需要开发人员逐个单元格写入数据，
略显笨拙。  
该项目提供类似于 easyPoi 的功能，利用 协程池快速的生成excel文件