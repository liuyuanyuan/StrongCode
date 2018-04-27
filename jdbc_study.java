
Oracle ref :  http://www.oracle.com/technetwork/java/overview-141217.html
JDBC中文教程： https://www.yiibai.com/jdbc/jdbc-driver-types.html


JDBC模型 — 深入理解JDBC设计思想（探究Class.forName("DBDriver")）：
参考: https://blog.csdn.net/daijin888888/article/details/50969621
ORACLE的SUN公司
首先，SUN公司定义了一个接口类：
/** 
 * @Description:sun公司定义的接口类 
 * 
 * @author:SUN 
 */  
public interface Connection {  
  
    public void f1();  
}  
同时，SUN公司定义了一个驱动管理类：
/** 
 * @Description:SUN公司定义的驱动管理类 
 * 
 * @author:SUN 
 */  
public class DriverManager {  
  
    public static Connection conn = null;  
  
    /** 
     * 注册连接 
     * 
     * @param connection 
     */  
    public static void registConnection(Connection connection) {  
        conn = connection;  
    }  
  
    /** 
     * 获取连接 
     * 
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


