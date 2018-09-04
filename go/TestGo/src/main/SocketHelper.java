

import java.io.BufferedReader;
import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.net.Socket;
import java.net.UnknownHostException;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;


public class SocketHelper 
{
	private final Logger logger = LoggerFactory.getLogger(getClass());
	
	public static void main(String[] args)
	{
		try
		{
			SocketHelper sh = new SocketHelper("192.168.100.170", 1208);
			//sh.sendFile("e:","123.txt","f:");
			
			String cmd = "ping";//"hello\n\r";
			System.out.println("cmd:" + cmd);
			boolean flag =  sh.executeCommand(cmd);
			System.out.println(flag);
			if(flag)
			{
				System.out.println("Return:"+ sh.getReturnMsg());		
			}	
			
		}catch(Exception ex)
		{
			System.err.println(ex.getMessage());
		}
	}

	private Socket socket;
    private StringBuilder str;
	public SocketHelper(String host,int port) throws Exception 
	{
		socket = null;				
		try 
		{
			//1.创建客户端Socket，指定服务器地址和端口
			socket = new Socket(host, port);
			
		} catch (UnknownHostException e) 
		{
			logger.error(e.getMessage());
			e.printStackTrace();
			throw new Exception(e);
		} catch (IOException e) 
		{
			logger.error(e.getMessage());
			e.printStackTrace();
			throw new Exception(e);
		}		
	}
	
	public boolean executeCommand(String cmd) throws Exception 
	{
		str = new StringBuilder();
		boolean success = false;
		
		OutputStream os = null;
		PrintWriter pw = null;
		InputStream is = null;
		BufferedReader br = null;
		try 
		{
			// 1.创建客户端Socket，指定服务器地址和端口
			//Socket socket = new Socket(host, 9999);

			// 2.获取输出流，向服务器端发送信息
			os = socket.getOutputStream();// 字节输出流
			pw = new PrintWriter(os);//将输出流包装为打印流
			pw.write(cmd);
			pw.flush();
			socket.shutdownOutput();//关闭输出流
			logger.info("shutdownOutput:" + socket.isOutputShutdown());

			// 3.获取输入流，并读取服务器端的响应信息
			is = socket.getInputStream();
			br = new BufferedReader(new InputStreamReader(is));
			
			String info = null;
			while ((info = br.readLine()) != null) 
			{
				str.append(info);
				logger.info(info);
			}
			
			success = true;			
		} catch (UnknownHostException e) 
		{
			logger.error(e.getMessage());
			e.printStackTrace();
			throw new Exception(e);
		} catch (IOException e) 
		{
			logger.error(e.getMessage());
			e.printStackTrace();
			throw new Exception(e);
		}finally
		{
			// 4.关闭资源
	        if(br!=null)
	        {
				try 
				{
					br.close();
				} catch (IOException e) 
				{
					e.printStackTrace();
				}
			}
			if (is != null) 
			{
				try
				{
					is.close();
				} catch (IOException e) 
				{
					e.printStackTrace();
				}
			}
			if (pw != null) 
			{
				pw.close();
			}
			if (os != null) 
			{
				try 
				{
					os.close();
				} catch (IOException e) 
				{
					e.printStackTrace();
				}
			}
			//socket.close();
		}
		logger.debug("Return:" + success);
		return success;
	}
	public String getReturnMsg()
	{
		return str.toString();
	}
	
	public void sendFile(String localPath, String fileName, String remotePath)
    {
		logger.debug("Enter:localPath=" + localPath + ",fileName=" + fileName
    			+ ",remotePath=" + remotePath );		
		DataOutputStream dos = null;
        DataInputStream dis = null;       
		try 
		{
			dos = new DataOutputStream(socket.getOutputStream());
			dis = new DataInputStream(socket.getInputStream());	    	
	    	dos.writeUTF("SEND");
	    	
	        File f = new File(localPath + File.separator + fileName);
	        if(!f.exists())
	        {
	            logger.debug("File not Exists, do nothing and return.");
	            dos.writeUTF("File not found");
	            return;
	        }
	        dos.writeUTF(remotePath +"/"+fileName);
	        
	        String msgFromServer = dis.readUTF();
	        if(msgFromServer.compareTo("File Already Exists") == 0)
	        {           
	            logger.debug("File Already Exists and we will Overwrite it.");
	        }
	        
	        logger.debug("Sending File ...");
	        FileInputStream fin=new FileInputStream(f);
	        int ch;
	        do
	        {
	            ch=fin.read();
	            dos.writeUTF(String.valueOf(ch));
	        }
	        while(ch!=-1);
	        fin.close();
	        logger.debug(dis.readUTF());
	        
	        //dos.writeUTF("DISCONNECT");
		} catch (Exception e) 
		{
			logger.error(e.getMessage());
			e.printStackTrace();
		} finally 
		{
			if (dis != null)
				try 
			    {
					dis.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			if (dos != null)
				try 
			    {
					dos.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			/*
			if (socket != null)
				try 
			    {
					socket.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			*/
		}
    }
   
	
	public void closeSocket()
	{
		if(socket != null)
		{
			try 
			{
				socket.close();
			} catch (IOException e) 
			{
				e.printStackTrace();
			}
		}
	}
	
	/*
	public void method()
	{
		try
		{
			Socket socket=new Socket("127.0.0.1",4700);
			//向本机的4700端口发出客户请求
			BufferedReader sin=new BufferedReader(new InputStreamReader(System.in));
			//由系统标准输入设备构造BufferedReader对象
			PrintWriter os=new PrintWriter(socket.getOutputStream());
			//由Socket对象得到输出流，并构造PrintWriter对象
			BufferedReader is=new BufferedReader(new InputStreamReader(socket.getInputStream()));
			//由Socket对象得到输入流，并构造相应的BufferedReader对象
			String readline;
			readline=sin.readLine(); //从系统标准输入读入一字符串
			while(!readline.equals("bye"))
			{
				//若从标准输入读入的字符串为 "bye"则停止循环
				os.println(readline);
				//将从系统标准输入读入的字符串输出到Server
				os.flush();
				//刷新输出流，使Server马上收到该字符串
				System.out.println("Client:"+readline);
				//在系统标准输出上打印读入的字符串
				System.out.println("Server:"+is.readLine());
				//从Server读入一字符串，并打印到标准输出上
				readline=sin.readLine(); //从系统标准输入读入一字符串
			} //继续循环
			os.close(); //关闭Socket输出流
			is.close(); //关闭Socket输入流
			socket.close(); //关闭Socket
		}catch(Exception e) 
		{
			System.out.println("Error"+e.getMessage());
		}
	}
	 
	 
	public static void main(String args[])
	{
		
		try
		{
			Socket socket=new Socket("127.0.0.1",9999);
			//向本机的4700端口发出客户请求
			
			BufferedReader sin=new BufferedReader(new InputStreamReader(System.in));
			//由系统标准输入设备构造BufferedReader对象
			
			PrintWriter os=new PrintWriter(socket.getOutputStream());
			//由Socket对象得到输出流，并构造PrintWriter对象
			
			BufferedReader is=new BufferedReader(new InputStreamReader(socket.getInputStream()));
			//由Socket对象得到输入流，并构造相应的BufferedReader对象
			
			String readline=sin.readLine(); 
			//从系统标准输入读入一字符串
			while(!readline.equals("bye"))
			{
				//若从标准输入读入的字符串为 "bye"则停止循环
				os.println(readline);
				//将从系统标准输入读入的字符串输出到Server
				os.flush();
				//刷新输出流，使Server马上收到该字符串
				System.out.println("Client:"+readline);
				//在系统标准输出上打印读入的字符串
				System.out.println("Server:"+is.readLine());
				//从Server读入一字符串，并打印到标准输出上
				readline=sin.readLine(); //从系统标准输入读入一字符串
			}
				//继续循环
				os.close(); //关闭Socket输出流
				is.close();//关闭Socket输入流
				socket.close(); //关闭Socket
			}catch(Exception e) 
	        {
				System.out.println(e.getMessage());//出错，则打印出错信息
			}
	}
	

	public void method1() 
	{
		Socket socket = null;
		try 
		{
			socket = new Socket("127.0.0.1", 9999);
			System.out.println("connected");

			//向服务器端第一次发送字符串
			OutputStream netOut = socket.getOutputStream();
			DataOutputStream doc = new DataOutputStream(netOut);
			DataInputStream in = new DataInputStream(socket.getInputStream());
			// 向服务器端第二次发送字符串
			doc.writeUTF("list\n");
			doc.flush();
			socket.shutdownOutput();	
			while(!socket.isOutputShutdown())
			{
				logger.info("wait");
			}
			System.out.println("sended list"+socket.isOutputShutdown());

			String res = in.readUTF();
			System.out.println("return:" + res);

			doc.writeUTF("bye\n");
			res = in.readUTF();
			System.out.println(res);

			// close all
			doc.close();
			in.close();
		} catch (UnknownHostException e) 
		{
			e.printStackTrace();
		} catch (IOException e) 
		{
			e.printStackTrace();
		} finally 
		{
			if (socket != null) 
			{
				try 
				{
					socket.close();
				} catch (IOException e) 
				{
					e.printStackTrace();
				}
			}
		}
	}
	*/

	
}
