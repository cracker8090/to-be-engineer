public class ChildClass implements Cloneable {
    public String name;
    public int age;

    @Override
    public Object clone(){
        try {
            return super.clone();
        } catch (CloneNotSupportedException ignore){
        }
        return null;
    }
}
