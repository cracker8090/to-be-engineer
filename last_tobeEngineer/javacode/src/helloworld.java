import java.util.logging.Logger;

public class helloworld {

    private static final String TAG = "helloworld";
    public static Logger log = Logger.getLogger(helloworld.class.toString());
    public static void main(String[] args) {
//        System.out.println("hello world");
//        System.out.println(args[0]);
//        System.out.println(args[1]);
        father fatherA = new father();
        fatherA.name = "zhang 3";
        fatherA.age = 26;
        fatherA.child = new ChildClass();
        fatherA.child.name = "xiao zhang 3";
        fatherA.child.age = 5;

        father fatherB = (father) fatherA.clone();
        log.info("fatherA=fatherB is "+(fatherA==fatherB));
        log.info("fatherA.hashCode="+fatherA.hashCode());
        log.info("fine:"+fatherA.name);
        log.info("fine:"+fatherA.age);
        log.info("fatherB.hashCode="+fatherB.hashCode());
        log.info("fine:"+fatherB.name);
        log.info("fine:"+fatherB.age);


    }


}




