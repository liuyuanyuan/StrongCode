package com.highgo.admin;

import org.eclipse.core.runtime.IConfigurationElement;
import org.jkiss.dbeaver.DBException;
import org.jkiss.dbeaver.model.DBPImage;
import org.jkiss.dbeaver.model.DBPObject;
import org.jkiss.dbeaver.model.struct.DBSObject;
import org.jkiss.dbeaver.registry.AbstractContextDescriptor;
import org.jkiss.dbeaver.registry.RegistryConstants;
import org.jkiss.dbeaver.tools.IExternalTool;
import org.jkiss.utils.CommonUtils;

public class AdvanceToolDescriptor extends AbstractContextDescriptor{
	private final String id;
	private String label;
	private String description;	
	private final DBPImage icon;
	
	private ObjectType toolType;
	private final boolean singleton;
	
	public AdvanceToolDescriptor(IConfigurationElement config)
	{
		super(config);
		this.id = config.getAttribute(RegistryConstants.ATTR_ID);
		this.label = config.getAttribute(RegistryConstants.ATTR_LABEL);
		this.description = config.getAttribute(RegistryConstants.ATTR_DESCRIPTION);		
		this.icon = iconToImage(config.getAttribute(RegistryConstants.ATTR_ICON));
		
		this.toolType = new ObjectType(config.getAttribute(RegistryConstants.ATTR_CLASS));
		this.singleton = CommonUtils.toBoolean(config.getAttribute(RegistryConstants.ATTR_SINGLETON));
	}

	public String getId() {
		return id;
	}

	public String getLabel() {
		return label;
	}

	public String getDescription() {
		return description;
	}

	public DBPImage getIcon() {
		return icon;
	}

	public boolean isSingleton() {
		return singleton;
	}

	@Override
	protected Object adaptType(DBPObject object) {
		if (object instanceof DBSObject) {
			return ((DBSObject) object).getDataSource();
		}
		return super.adaptType(object);
	}

	public IExternalTool createTool() throws DBException {
		return toolType.createInstance(IExternalTool.class);
	}
}
