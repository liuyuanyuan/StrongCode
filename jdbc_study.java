=pgjdbc dev and build jar=
source: https://github.com/pgjdbc/pgjdbc
dev: eclipse import as maven project from .\pgjdbc-master\pgjdbc (ignore pom.xml error)
build jar: eclipse run as - maven build
           eclipse run as - run configuration:(Goals:jar:jar --ref from pom.xml maven-jar-plugin) (skip test:check) - Apply and Run
           postgresql-42.2.6-SNAPSHOT.jar: will generated in .\pgjdbc-master\pgjdbc\target
skip test: because some url in java for test can not pass.

source main:
     PGStream: a socket to connect to pg db;
  
     PGCallableStatement: callableStatement() for plpgsql function or procedure
     
     PGPreparedStatement: preparedStatement() for sql query
     QueryExecutorImpl: to send query by different preferQueryMode
     QueryExecutorBase:
 


==================================================================================================================
 
0.Oracle JDBC ref : http://www.oracle.com/technetwork/java/overview-141217.html
The JDBC API is the industry standard for database-independent connectivity between the Java programming language and a wide range of databases. The JDBC API provides a call-level API for SQL-based database access. JDBC technology allows you to use the Java programming language to exploit "Write Once, Run Anywhere" capabilities for applications that require access to enterprise data.
 
JDBC API Overview

The JDBC API makes it possible to do three things:
Establish a connection with a database or access any tabular data source
Send SQL statements
Process the results

------------------------------------------------------------------
1.JDBC模型 — 深入理解JDBC设计思想（探究Class.forName("DBDriver")）：
参考: https://blog.csdn.net/daijin888888/article/details/50969621
ORACLE的SUN公司
首先，SUN公司定义了一个接口类：
/** 
 * @Description: SUN公司定义的接口类 
 * 
 * @author:SUN 
 */  
public interface Connection {    
    public void f1();  
}  
同时，SUN公司定义了一个驱动管理类：
/** 
 * @Description: SUN公司定义的驱动管理类 
 * 
 * @author:SUN 
 */  
public class DriverManager { 
    public static Connection conn = null; 
    /** 
     * 注册连接 
     * @param connection 
     */  
    public static void registConnection(Connection connection) {  
        conn = connection;  
    }  
  
    /** 
     * 获取连接  
     * @return 
     */  
    public static Connection getConnection() {  
        return conn;  
    }  
} 

数据库厂商：（如Oracle、MySQL等）
各数据库厂商实现SUN制定的接口标准，如：Oracle

/** 
 * @Description:Oracle数据库厂商实现的接口类 
 * 
 * @author:Oracle 
 */  
public class ConnectionOracleImpl implements Connection {    
    @Override  
    public void f1() {  
        // 这里实现Oracle操作数据库的具体方法，封装在.jar文件中，供程序员调用  
        System.out.println("Oracle的f1()方法实现");  
    }    
}  

或者MySQL
/** 
 * @Description:MySQL数据库厂商实现的接口类 
 * 
 * @author:MySQL 
 */  
public class ConnectionMySQLImpl implements Connection { 
    @Override  
    public void f1() {  
        //这里实现MySQL操作数据库的具体方法，封装在.jar文件中，供程序员调用  
        System.out.println("MySQL的f1()方法实现");  
    }
}  

同时，数据库厂商定制自己的驱动类，并通过静态代码块，确保该驱动类被加载时，SUN的驱动管理类可以注册厂商的数据库连接，如：
Oracle
/** 
 * @Description:Oracle厂商制定的驱动类 
 * 
 * @author:Oracle 
 */  
public class OracleDriver {  
    static {  
        DriverManager.registConnection(new ConnectionOracleImpl());  
    }  
}  

或者MySQL
/** 
 * @Description:MySQL厂商制定的驱动类 
 * 
 * @author:MySQL 
 */  
public class MySQLDriver {  
    static {  
        DriverManager.registConnection(new ConnectionMySQLImpl());  
    }  
}  

开发者：
开发人员引入某个数据库厂商的驱动文件，即可调用内部方法操作数据库，如：
/** 
 * @Description:开发人员 
 * 
 * @author:me 
 */  
public class Test2 {  
  
    public static void main(String[] args) throws Exception {  
        Class.forName("test.OracleDriver");// 虚拟机根据类名找到字节码文件  
        Connection con = DriverManager.getConnection();  
        con.f1();  
  
        Class.forName("test.MySQLDriver");// 虚拟机根据类名找到字节码文件  
        Connection con2 = DriverManager.getConnection();  
        con2.f1();  
    }  
}  

运行结果：
Oracle的f1()方法实现
MySQL的f1()方法实现

---------------------------------------------------------
2.JDBC要点总结、SQL注入示例（Statement和PreparedStatement）
https://blog.csdn.net/daijin888888/article/details/50965232

使用原有Statement有以下问题：
--容易遭受注入式攻击
--拼写SQL繁琐和麻烦
通过PrepareStatement可以解决上述问题，因为：
PrepareStatement是一个预编译的Statement，将带占位符？的SQL语句发送给数据库后，SQL语句不会立即执行，数据库会生成一个执行计划，此时SQL语句结构已确定，不可更改注入，然后利用setXXX()方法给sql的?设置参数值，传参后即执行计划，返回结果集。另外，由于执行计划已生成，只要传入参数就可执行计划，这在大批量存入数据时，编码更简单，效率更高。

PrepareStatement使用步骤：
--编写带?号的sql
--利用con.prepareStatement(sql);方法获取PrepareStatement对象
--利用setXXX()方法给sql的?设置参数值
--调用无参的executeUpdate()或executeQuery()执行sql.

------------------------------------------------------------------
3.JDBC中文教程： https://www.yiibai.com/jdbc/jdbc-driver-types.html
2.1 SQLException中的方法
一个SQLException类既可以发生在驱动程序和数据库中。当这样的异常时，SQLException类型的对象将被传递到catch子句。
通过SQLException对象有以下几种方法可用于获取更多的关于异常的信息：

方法	描述
getErrorCode()	获取与异常关联的错误代码
getMessage()	获取JDBC驱动程序的错误处理错误消息，或获取Oracle错误代码和数据库的错误消息。
getSQLState()	获取XOPEN SQLSTATE字符串。对于JDBC驱动程序错误，从该方法返回的可能是无用的信息。对于一个数据库错误，返回一个五位的XOPEN SQLSTATE代码。这种方法可以返回null。
getNextException()	获取异常链中的下一个Exception对象
printStackTrace()	打印当前的异常，或也可以抛出，并回溯到标准错误流
printStackTrace(PrintStream s)	打印此抛出对象及其回溯到指定的打印流
printStackTrace(PrintWriter w)	打印此抛出对象及其回溯到指定打印写入流
通过利用从Exception对象提供的信息，可以捕获一个异常，并适当地继续运行程序。这是一个try块中的一般形式：
try {
   // Your risky code goes between these curly braces!!!
}catch(Exception ex) {
   // Your exception handling code goes between these 
   // curly braces, similar to the exception clause 
   // in a PL/SQL block.
}finally {
   // Your must-always-be-executed code goes between these 
   // curly braces. Like closing database connection.
}

2.2 JDBC数据类型
下表列出了默认的JDBC数据类型与Java数据类型转换，当使用PreparedStatement或CallableStatement对象时可调用setXXX()方法或ResultSet.updateXXX()方法。
SQL	JDBC/Java	setXXX	updateXXX
VARCHAR	java.lang.String	setString	updateString
CHAR	java.lang.String	setString	updateString
LONGVARCHAR	java.lang.String	setString	updateString
BIT	boolean	setBoolean	updateBoolean
NUMERIC	java.math.BigDecimal	setBigDecimal	updateBigDecimal
TINYINT	byte	setByte	updateByte
SMALLINT	short	setShort	updateShort
INTEGER	int	setInt	updateInt
BIGINT	long	setLong	updateLong
REAL	float	setFloat	updateFloat
FLOAT	float	setFloat	updateFloat
DOUBLE	double	setDouble	updateDouble
VARBINARY	byte[ ]	setBytes	updateBytes
BINARY	byte[ ]	setBytes	updateBytes
DATE	java.sql.Date	setDate	updateDate
TIME	java.sql.Time	setTime	updateTime
TIMESTAMP	java.sql.Timestamp	setTimestamp	updateTimestamp
CLOB	java.sql.Clob	setClob	updateClob
BLOB	java.sql.Blob	setBlob	updateBlob
ARRAY	java.sql.Array	setARRAY	updateARRAY
REF	java.sql.Ref	SetRef	updateRef
STRUCT	java.sql.Struct	SetStruct	updateStruct
在JDBC3.0中增强支持BLOB，CLOB，ARRAY，REF等数据类型。ResultSet对象可调用UPDATEBLOB()，updateCLOB()，updateArray()和updateRef()方法，使您可以在数据库服务器上直接操作相应的数据。

对于setXXX()和updateXXX()方法，可以转换成特定的Java类型到特定的JDBC数据类型。而使用setObject()和updateObject()方法，几乎所有的Java类型映射到JDBC数据类型。

ResultSet对象提供相应的getXXX()方法为每个数据类型来检索列值。每一种类型方法，可以使用与列名或由列的序号位置来获取列的数据。

SQL	JDBC/Java	setXXX	getXXX
VARCHAR	java.lang.String	setString	getString
CHAR	java.lang.String	setString	getString
LONGVARCHAR	java.lang.String	setString	getString
BIT	boolean	setBoolean	getBoolean
NUMERIC	java.math.BigDecimal	setBigDecimal	getBigDecimal
TINYINT	byte	setByte	getByte
SMALLINT	short	setShort	getShort
INTEGER	int	setInt	getInt
BIGINT	long	setLong	getLong
REAL	float	setFloat	getFloat
FLOAT	float	setFloat	getFloat
DOUBLE	double	setDouble	getDouble
VARBINARY	byte[ ]	setBytes	getBytes
BINARY	byte[ ]	setBytes	getBytes
DATE	java.sql.Date	setDate	getDate
TIME	java.sql.Time	setTime	getTime
TIMESTAMP	java.sql.Timestamp	setTimestamp	getTimestamp
CLOB	java.sql.Clob	setClob	getClob
BLOB	java.sql.Blob	setBlob	getBlob
ARRAY	java.sql.Array	setARRAY	getARRAY
REF	java.sql.Ref	SetRef	getRef
STRUCT	java.sql.Struct	SetStruct	getStruct

2.3 JDBC批量处理
批处理允许执行一个批处理组相关的SQL语句，并将其一次提交到数据库中执行。当几个SQL语句一次发送到数据库中时，可以减少通信开销，从而提高性能。

JDBC驱动程序不支持此功能。您应该使用DatabaseMetaData.supportsBatchUpdates()方法来确定目标数据库支持批量更新处理。如果JDBC驱动程序支持此功能，则该方法返回true。
addBatch()方法是PreparedStatement和CallableStatementis类中用于添加单个语句的批处理的声明。 executeBatch()将开始将所有语句组合到一起并执行。

executeBatch()将返回一个整数数组，每个数组元素的表示为相应的更新语句的更新计数。

添加语句进行批处理时，可以使用clearBatch()方法删除它们。此方法将删除addBatch()方法添加的所有语句。但是不能有选择性地选择某个语句来删除。

2.4
JDBC数据流
PreparedStatement对象有能力使用提供参数数据的输入和输出流。这使您可以将整个文件到数据库中，可容纳较大的值，如CLOB和BLOB数据类型的列。

有下列方法可用于流数据：
setAsciiStream(): 此方法用于提供大的ASCII数据值。
setCharacterStream(): 此方法用于提供大的UNICODE数据值。
setBinaryStream(): 使用此方法用于提供大的二进制数据值。
setXXXStream()方法需要一个额外的参数，文件大小(除了参数占位符)。此参数通知应发送多少数据到数据库来使用流的驱动程序。
对于一个详细的关于所有这些概念，这里只是一个简单的入门教程，还需要读者去学习完整的教程，有关JDBC的后续教程，可以从左侧文章中了解和学习。


