项目管理工具--maven
--------------------
maven是用于管理包依赖和生成过程的配置。使用一个叫做pom.xml的文件来管理

maven建立了一个本地仓库和中央仓库来管理所有依赖，相关仓库保存位置:`$user.home/.m2/repository`

# 安装

* 下载地址:[点击此处](http://maven.apache.org/download.html)
* 添加环境变量:`%安装位置%\bin`
* 验证是否安装成功:`mvn -version`

# 使用mvn

* 创建一个mvn项目的根目录（不同项目放在不同目录）
* 使用命令`mvn archetype:generate`初始化一个mvn项目（第一次运行时，会下载很多依赖包），初始化时，会要求填写一些参数:
  * groupId:项目的唯一Id
  * artifactId:生成的项目目录的名称
  * version: 项目版本号
  * package: 项目source文件夹下最顶层的包名

# pom.xml配置文件解析

`pom.xml`文件是mvn管理一个项目的核心文件。
````
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <!-- 指定了当前POM的版本 -->
    <modelVersion>4.0.0</modelVersion>

	<!-- 公司或者组织的唯一标志，并且配置时生成的路径也是由此生成， 如com.companyname.project-group，maven会将该项目打成的jar包放本地路径：/com/companyname/project-group -->
	<groupId>com.sogou.hi</groupId>
    <!-- 项目的唯一ID，一个groupId下面可能多个项目，就是靠artifactId来区分的 -->
	<artifactId>hi</artifactId>
	<!-- 大版本号.分支版本号.小版本号  snapshot 快照  alpha内测  beta公测  release稳定版  GA正式发布版本 -->
	<version>0.0.1-SNAPSHOT</version>
	<!--项目产生的构件类型，例如jar、war、ear、pom。插件可以创建他们自己的构件类型 -->
	<packaging>pom</packaging><!--指定打包方式-->

	<!--项目的名称, Maven产生的文档用 -->
	<name>ml-parent</name>
	<!--项目主页的URL, Maven产生的文档用 -->
	<url>http://maven.apache.org</url>
    <!-- 项目的详细描述, Maven 产生的文档用。 当这个元素能够用HTML格式描述时（例如，CDATA中的文本会被解析器忽略，就可以包含HTML标 
        签）， 不鼓励使用纯文本描述。如果你需要修改产生的web站点的索引页面，你应该修改你自己的索引页文件，而不是调整这里的文档。 -->
    <description>A maven project to study maven.</description>

	<!--其他用于解析的自定义属性-->
   <properties>
     <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
     <junit.version>3.8.1</junit.version>
   </properties>

    <!-- 指定当前的依赖配置，是一个数组 -->
	<dependencyManagement>
		<dependencies>
		    <dependency>
		      <groupId>junit</groupId> <!--依赖的项目名 -->
		      <artifactId>junit</artifactId> <!--依赖的项目模块名-->
		      <version>3.8.1</version> <!--使用的项目版本-->
		      <scope>test</scope><!-- scope 属性决定该依赖项目在什么阶段，test表示该项目只在测试代码中依赖 具体见  http://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html#Dependency_Scope -->
		    </dependency>
		     <dependency>
		      <groupId>com.sogou.ml</groupId>
		      <artifactId>ml-b</artifactId>
		      <version>0.0.1-SNAPSHOT</version>
		      <exclusions> <!-- 排除传递关系的依赖。 例如 ml-c 依赖 ml-b，ml-b依赖 ml-a ，那么我们会发现maven让ml-c同时依赖了a和b，通过这个属性可以排除c对a的依赖 -->
		          <exclusion>
		              <groupId>com.sogou.ml</groupId>
		              <artifactId>ml-a</artifactId>
		          </exclusion>
		      </exclusions>
		    </dependency>
		  </dependencies>
	
	</dependencyManagement>

	<!--指定项目的生成过程-->
	<build>
        <!-- 插件列表 -->
        <plugins>
            <!-- 打包源码插件 -->
            <plugin>
                <!-- 插件项目坐标 -->
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-source-plugin</artifactId>
                <version>2.4</version>
                <!-- 在什么阶段执行 -->
                <executions>
                    <execution>
                        <phase>package</phase>
                        <goals>
                            <goal>jar-no-fork</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
        </plugins>
        <!--该元素设置了项目源码目录，当构建项目的时候，构建系统会编译目录里的源码。该路径是相对于pom.xml的相对路径。 -->
        <sourceDirectory />
        <!--该元素设置了项目脚本源码目录，该目录和源码目录不同：绝大多数情况下，该目录下的内容 会被拷贝到输出目录(因为脚本是被解释的，而不是被编译的)。 -->
        <scriptSourceDirectory />
        <!--该元素设置了项目单元测试使用的源码目录，当测试项目的时候，构建系统会编译目录里的源码。该路径是相对于pom.xml的相对路径。 -->
        <testSourceDirectory />
        <!--被编译过的应用程序class文件存放的目录。 -->
        <outputDirectory />
        <!--被编译过的测试class文件存放的目录。 -->
        <testOutputDirectory />

    </build>
</project>
````

## 使用maven打包多个项目
添加modules属性，这个属性也是一个数组，里边有一些项目文件夹的路径，编译打包这个pom项目的时候，会将modules里边的项目都编译打包。
````
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.sogou.ml</groupId>
  <artifactId>ml-all</artifactId>
  <version>0.0.1-SNAPSHOT</version>
  <packaging>pom</packaging>

  <name>ml-all</name>
  <url>http://maven.apache.org</url>
  <modules>
      <module><!--项目1名-->
          ../ml-a
      </module>
      <module><!--项目2名-->
          ../ml-b
      </module>
      <module>
          ../ml-c
      </module>
  </modules>
  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>
</project>
````

# 相关命令

# 参考资料
* [maven 菜鸟教程](https://www.runoob.com/maven/maven-tutorial.html)