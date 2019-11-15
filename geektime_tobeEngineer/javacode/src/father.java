public class father implements Cloneable {
    public String name;
    public int age;
    public ChildClass child;

    @Override
    public Object clone(){
        try {
            father fatherclone = (father) super.clone();
            fatherclone.child = (ChildClass)this.child.clone();
            return fatherclone;
        } catch (CloneNotSupportedException ignore){
        }
        return null;
    }



}
