package com.highgo.admin;

import java.util.ArrayList;
import java.util.List;

import org.eclipse.core.runtime.IConfigurationElement;
import org.eclipse.core.runtime.IExtensionRegistry;
import org.eclipse.core.runtime.Platform;
import org.jkiss.dbeaver.Log;
/*
 * access extension
*/
public class AdvanceToolsRegistry {
	private Log log = Log.getLog(AdvanceToolsContributor.class);	
	
	
	private static AdvanceToolsRegistry instance = null;
	public synchronized static AdvanceToolsRegistry getInstance() 
	{
		if (instance == null) 
		{
			instance = new AdvanceToolsRegistry();
			instance.loadExtensions(Platform.getExtensionRegistry());
		}
		return instance;
	}

	public static final String EXTENSION_POINT_ID = "com.highgo.admin.advance.tools";
	static final String TAG_TOOL = "tool";
	private final List<AdvanceToolDescriptor> tools = new ArrayList<>();	
	private void loadExtensions(IExtensionRegistry registery) 
	{	
		log.debug("AdvanceToolsRegistry.loadExtensions: enter");
		
		IConfigurationElement[] extConfigs = registery.getConfigurationElementsFor(EXTENSION_POINT_ID);//
		log.debug("extConfigs size=" + extConfigs.length);
		for (IConfigurationElement toolsElement : extConfigs) 
		{
			log.debug("tools element=" + toolsElement.getValue());
			log.debug("tool size" + toolsElement.getChildren(TAG_TOOL).length);
			for (IConfigurationElement toolElement : toolsElement.getChildren(TAG_TOOL))//
			{
				log.debug("tool element=" + toolElement.getValue());
				this.tools.add(new AdvanceToolDescriptor(toolElement));
			}
		}
	}

	public List<AdvanceToolDescriptor> getTools() {
		return tools;
	}

	public void dispose() {
		tools.clear();
	}
	
}
