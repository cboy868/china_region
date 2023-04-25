china_region 中国省市地区的数据采集
_____


数据主页:[http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/13.html](http://www.stats.gov.cn/sj/tjbz/tjyqhdmhcxhfdm/2022/13.html)  


使用
_____

```
git clone https://github.com/wecatch/china_regions.git
cd china_regions
go run main.go
```


目前需要连接redis


行政级别顺序
_____
province -> city --> country --> town --> village  
省->市(市辖区)->县(区、市)->镇(街道)->村(居委会)  


ToDo
_____
1.并发  
2.生成多表数据插入sql 省，市(市辖区)，县(区、市)，镇(街道)，村(居委会)  
3.生成单表数据插入sql  
