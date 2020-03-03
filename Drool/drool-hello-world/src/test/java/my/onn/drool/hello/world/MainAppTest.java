package my.onn.drool.hello.world;

import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.Reader;
import org.drools.compiler.compiler.DroolsParserException;
import org.drools.compiler.compiler.PackageBuilder;
import org.drools.core.RuleBase;
import org.drools.core.RuleBaseFactory;
import org.drools.core.WorkingMemory;

public class MainAppTest {

    public static void main(String[] args) throws DroolsParserException, IOException {
        MainAppTest droolsTest = new MainAppTest();
        droolsTest.executeDrools();
    }

    public void executeDrools() throws DroolsParserException, IOException {
        PackageBuilder packageBuilder = new PackageBuilder();
        
        String ruleFile = "/com/rule/ProductRule.drl";
        InputStream resourceAsStream = getClass().getResourceAsStream(ruleFile);
        
        Reader reader = new InputStreamReader(resourceAsStream);
        packageBuilder.addPackageFromDrl(reader);
        org.drools.core.rule.Package rulesPackage = packageBuilder.getPackage();
        RuleBase ruleBase = RuleBaseFactory.newRuleBase();
        ruleBase.addPackage(rulesPackage);
        
        WorkingMemory workingMemory = ruleBase.newStatefulSession();
        
        Product product = new Product();
        product.setType("diamond");
        
        workingMemory.insert(product);
        workingMemory.fireAllRules();
        
        System.out.println(String.format("The discount for product %s is %d", product.getType(), product.getDiscount()));
    }

    
}
