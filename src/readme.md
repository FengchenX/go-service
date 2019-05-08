restful文档: https://blog.csdn.net/qq_41606973/article/details/86352787

URI 表示资源，资源一般对应服务器端领域模型中的实体类。
不用大写；
用中杠-不用下杠_;
URI中的名词表示资源集合，使用复数形式。

资源集合：
/zoos //所有动物园
/zoos/1/animals //id为1的动物园中的所有动物

单个资源：
/zoos/1 //id为1的动物园
/zoos/1;2;3 //id为1，2，3的动物园


/在url中表达层级，用于按实体关联关系进行对象导航，一般根据id导航。
过深的导航容易导致url膨胀，不易维护，如 GET /zoos/1/areas/3/animals/4，尽量使用查询参数代替路径中的实体导航，
如GET /animals?zoo=1&area=3；

对Composite资源的访问
一个常见的例子是 User — Address，Address是对User表中zipCode/country/city三个字段的简单抽象，无法独立于User存在。
必须通过User索引到Address：GET /user/1/addresses

response 的 body 直接就是数据，不要做多余的包装。错误示例：
{
    "success":true,
    "data":{"id":1,"name":"xiaotuan"},
}
·	response 格式
GET	       单个对象、集合
POST	   新增成功的对象
PUT/PATCH  更新成功的对象
DELETE	   空

entity 与数据库表一一对应
dto 中间转换层
vo 与界面一一对应

